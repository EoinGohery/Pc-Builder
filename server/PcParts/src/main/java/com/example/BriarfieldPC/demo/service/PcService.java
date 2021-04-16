package com.example.BriarfieldPC.demo.service;

import com.example.BriarfieldPC.demo.entity.*;
import com.example.BriarfieldPC.demo.repository.*;
import lombok.AllArgsConstructor;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@AllArgsConstructor
public class PcService {
    @Autowired
    private CpuRepositroy cpuRepository;
    @Autowired
    private GpuRepository gpuRepository;
    @Autowired
    private DriveRepository driveRepository;
    @Autowired
    private MotherboardRepository motherboardRepository;
    @Autowired
    private PsuRepository psuRepository;
    @Autowired
    private RamRepository ramRepository;

    public List<Cpu> getAllCpu(){
        return cpuRepository.findAll();
    }
    public List<Gpu> getAllGpu(){
        return gpuRepository.findAll();
    }
    public List<Drive> getAllDrivers(){
        return driveRepository.findAll();
    }
    public List<Motherboard> getAllMotherboards(){
        return motherboardRepository.findAll();
    }
    public List<PSU> getAllPSU(){
        return psuRepository.findAll();
    }
    public List<Ram> getAllRam(){
        return ramRepository.findAll();
    }
}
